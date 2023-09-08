Get-Help "about*" -Full

Get-Command "*update*"

Update-Help -UICulture en-US

Clear-Host


[int]$num = Read-Host "Enter a number"
if (-not ($num -le 100)) {
    Write-Output "Entered number $num is NOT less than 100"
}