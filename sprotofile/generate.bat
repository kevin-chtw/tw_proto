@echo off
REM Windows batch script for generating Go protobuf files

REM Check if protoc is available
where protoc >nul 2>nul
if %errorlevel% neq 0 (
    echo Error: protoc not found. Please install Protocol Buffers compiler.
    echo Download from: https://github.com/protocolbuffers/protobuf/releases
    pause
    exit /b 1
)

REM Generate Go files
protoc -I=. --go_out=../sproto *.proto

if %errorlevel% equ 0 (
    echo Go protobuf files generated successfully.
) else (
    echo Error: Failed to generate Go protobuf files.
    pause
    exit /b 1
)

pause