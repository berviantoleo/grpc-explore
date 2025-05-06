RET=1
until [ ${RET} -eq 0 ]; do
    dotnet build -c Release
    RET=$?
    sleep 10
done