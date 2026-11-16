# Abre o tunel SSH para o Postgres da Hostinger e roda a aplicacao.
# Uso: .\iniciar.ps1   (pede a senha do banco; ela nao fica salva em arquivo)

$tunel = Start-Process ssh -ArgumentList "-N","-L","5433:127.0.0.1:5432","root@76.13.163.235" -WindowStyle Hidden -PassThru
Start-Sleep -Seconds 3

$senha = Read-Host "Senha do banco (usuario wesleyoliveira)" -AsSecureString
$env:DB_PASSWORD = [System.Net.NetworkCredential]::new("", $senha).Password

try {
    Write-Host "Abra http://localhost:8000/  (Ctrl+C para parar)"
    go run main.go
}
finally {
    Stop-Process -Id $tunel.Id -ErrorAction SilentlyContinue
}
