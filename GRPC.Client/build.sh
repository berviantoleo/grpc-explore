RET=1
until [ ${RET} -eq 0 ]; do
    dotnet build
    RET=$?
    sleep 10
done
