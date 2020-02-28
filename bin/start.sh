#!/bin/bash
pid=$(ps -ef | grep -w gzh | grep -v grep | awk '{print $2}')
if [[ ${pid} -gt 0 ]]; then
  kill -9 ${pid}
fi
cd bin
nohup ./gzh &