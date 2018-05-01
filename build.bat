call cd "source\frontend"
call npm install 
call npm run build 
call cd "..\backend"
call go build -x 
call cd "..\..\"
call move "source\backend\backend.exe" "dist\server.exe"
call cd "dist"
call server.exe