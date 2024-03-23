build:
	dotnet build

rebuild:
	dotnet build --no-incremental

publish:
	dotnet publish -c Release -o ./publish