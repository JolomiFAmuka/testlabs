
#configure mail service provider

$SMTPServer = "smtp.gmail.com"
$SMTPPort = "587"
$Username = "jolomi.amuka@gmail.com"
$Password = "liafscmuhansxwaj"

#Add receipient email address
$to = "j.amuka@yahoo.co.uk"
$cc = "ehimja@gmail.com"

$subject = "test autogenerate email"

$body = "
        Dear Mail Receiver,

        This is a powershell script generated email.
        Please do not reply.

        Information: This is then important information

        Regards,
        Team XYZ

        "

$message = New-Object System.Net.Mail.MailMessage
$message.Subject = $subject
$message.Body = $body
$message.To.Add($to)
$message.CC.add($cc)
$message.From = $Username

$message

$smtp = New-Object System.Net.Mail.SmtpClient($SMTPServer,$SMTPPort)
$smtp.EnableSsl = $true
$smtp.Credentials = New-Object System.Net.NetworkCredential($Username,$Password)

$smtp.Send($message)