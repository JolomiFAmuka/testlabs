$Global:g_vcar = "this is a global var"

$loc_var = "I'm a local var"

$cars = "toyota","range","benz"

foreach ($car in $cars) {
    Write-Output $car
}

$cars | gm

$array = "student1","student2","student3","student4","student5"

foreach ($item in $array) {
    Write-Output $item
}

$student_list = New-Object System.Collections.ArrayList

$student_list.Add("Male_Student")
$student_list.AddRange( ("Male_Student2","Female_Student3") )

$student_list | Get-Member


[String]$myName = Read-Host "Tell me your name"
$count = 0

while ($myName -ne "jolomi") {
    Write-Output "strange response...please try again"
    [String]$myName = Read-Host "Tell me your name"
    $count += 1
    Write-Output $count
    if ($count -eq 10) {
        break
    }
}

$i = 5
do {
    Write-Output "ciurrent count in the matter is $i"
    $i--
} while (
    $i -gt 0
)

function sayHi {
    Write-Output "Hello there"
}

sayHi
sayHi
sayHi
sayHi

function greeting {
    param (
        [String]$name,
        [int]$age
    )
    Write-Output "Hello there, my name is $name and i am $age years old"
}

# &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
# &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
$peeop = @{
    "mike" = @{"address" = "winnipeg drive"; "spouse" = "mitchell"; "age" = 51};
    "tony" = @{"address" = "chancellor drive"; "spouse" = "tania"; "age" = 41};
    "kirbie" = @{"address" = "lake drive"; "spouse" = "angela"; "age" = 28};
}

#sorting through the hash with foreach loop

 foreach ($person in $peeop.keys) {

    Write-Output "$person is $($peeop[$person]["age"])yrs old and spouse is $($peeop[$person]["spouse"]). They live at $($peeop[$person]["address"])"
 }

 #sorting through the hash foreah loop inside a function

 function vertoli($array) {
    foreach ($record in $array.keys) {

        Write-Output "$record is $($array[$record]["age"])yrs old and spouse is $($array[$record]["spouse"]). They live at $($array[$record]["address"])"
     }

 }

 vertoli($peeop)

 &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
 &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

 function myIntro {
    param (
        [String]$name,
        [String]$profession,
        [int]$age
    )
    Write-Output "Hello there i am $name, i work as a $profession and i'm $age yrs old"
 }


 gps | sort -descending cpu | select -first 10
