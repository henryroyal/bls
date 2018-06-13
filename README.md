
#for go in $(ls); do mv $go $(echo ${go}|sed s/.go/.sql/) done