#!/bin/bash

# 包路径
packagePath=""
# go的包名
packageName=""
# 服务名
serviceName=""
serviceFullName=""
packageService="service"

# 方法计数
funcInfoCount=0
# 所有方法的定义参数
# [方法名、请求参数、响应 ... ]
funcInfos=()

function toFirstLetterUpper(){
  str=$1
  firstLetter=${str:0:1}
  otherLetter=${str:1}
  firstLetter=$(echo $firstLetter | tr '[a-z]' '[A-Z]')
  result=$firstLetter$otherLetter
}

# 解析服务定义的方法
function parseRpcFunc() {
    funcDefine=$1
    # echo "解析方法：${funcDefine}"

    tmp=$(echo ${funcDefine} | sed -r 's/^rpc([a-zA-Z0-9_]+)\(([a-zA-Z0-9_]+)\)returns\(([a-zA-Z0-9_]+)\)(.*)/\1 \2 \3/')
    # echo "方法定义：${tmp}"

    # 方法名字
    tmpStr=$( echo ${tmp} | cut -d ' ' -f 1 )
    toFirstLetterUpper ${tmpStr}
    funcInfos[${funcInfoCount}]=${result}
    ((funcInfoCount++))

    # 请求参数
    tmpStr=$( echo ${tmp} | cut -d ' ' -f 2 )
    toFirstLetterUpper ${tmpStr}
    funcInfos[${funcInfoCount}]=${result}
    ((funcInfoCount++))

    # 响应
    tmpStr=$( echo ${tmp} | cut -d ' ' -f 3 )
    toFirstLetterUpper ${tmpStr}
    funcInfos[${funcInfoCount}]=${result}
    ((funcInfoCount++))
}


# 解析proto文件
# 获取包名、服务名、服务定义的方法
function parseProto() {

    packagePath="";packageName="";serviceName="";funcInfoCount=0;funcInfos=()

    if [ -e ${filePath} ]; then
        echo "========================================================="
        echo "正在解析：${filePath}"
    else
        echo "文件不存在：${filePath}"
        return 1
    fi

    # 临时变量：标识服务定义的开始
    serviceBegin=false
    # 临时变量：当前正在解析的方法
    rpcFunctionDefine=""
    # 临时变量：标识多行注释的开始
    commentBegin=false

    IFSBak=${IFS}
    IFS=$'\n'
    for line in $(cat ${filePath})
    do
        line=${line%//*}
        # 不处理空行或注释
        if [[ ${line} =~ ^[[:space:]]*$ || ${line} =~ // || ${commentBegin} == true ]]; then
            continue
        elif [[ ${line} =~ /\* ]]; then
            commentBegin=true
            continue
        elif [[ ${line} =~ \*/ ]]; then
            commentBegin=false
            continue
        fi

        if [[ -z "${packageName}" && ${line} =~ go_package ]]; then
            # 解析包名和路径
            tmp=$(echo ${line} | cut -d '"' -f 2)
            tmp=${tmp#*./}
            if [[ ${tmp} =~ ^/ ]]; then
                tmp=${tmp#*/}
            fi

            packagePath=$(echo ${tmp} | cut -d ';' -f 1)
            packageName=$(echo ${tmp} | cut -d ';' -f 2)
            if [ -z ${packageName} ]; then
                packageName=${packagePath%/*}
            fi

            echo "包路径：${packagePath} 包名：${packageName}"
        elif [[ -z "${serviceName}" && ${line} =~ ^service ]]; then
            # 解析服务名
            tmp=${line%%[{]*}
            serviceName=${tmp:8}
            serviceName=${serviceName%% *}
            toFirstLetterUpper ${serviceName}
            serviceName=${result}
            serviceFullName=${packageName}_${serviceName}
            echo "服务名：${serviceName}"
            serviceBegin=true
        elif [ ${serviceBegin} = true ]; then
            if [[ ${line} =~ [[:space:]]rpc[[:space:]] ]]; then
                if [[ -n "${rpcFunctionDefine}" ]]; then
                    # 上一个方法定义结束
                    rpcFunctionDefine=${rpcFunctionDefine//[[:space:]]/}
                    # echo "rpc方法定义被rpc结束： ${rpcFunctionDefine}"
                    parseRpcFunc ${rpcFunctionDefine}
                fi
                # 方法定义的起始行
                rpcFunctionDefine=${line}
                # echo "rpc方法定义开始： ${rpcFunctionDefine}"
            elif [[ ${line} =~ ^} ]]; then
                serviceBegin=false
                rpcFunctionDefine=${rpcFunctionDefine//[[:space:]]/}
                # echo "服务定义结束： ${rpcFunctionDefine}"
                parseRpcFunc ${rpcFunctionDefine}
                rpcFunctionDefine=""
            elif [[ -n "${rpcFunctionDefine}" ]]; then
                # 方法定义开始了但是还未结束
                rpcFunctionDefine=${rpcFunctionDefine}${line}
                # echo "rpc方法定义中：${rpcFunctionDefine}"
            fi
        fi
    done

    IFS=${IFSBak}

    if [ -z ${packageName} ]; then
        echo "没找到包名：${filePath}"
        return 1
    fi

    return 0
}