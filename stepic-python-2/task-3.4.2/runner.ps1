#$urls = Get-Content urls.html
#foreach ($url in $urls) {
##    Write-Host $url
#python -m venv C:\Projects\myzed\venv\

Get-Content "./objects.json" | c:\python39\python.exe task-3.4.2.py

