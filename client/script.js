var element = document.querySelector('#result')

const date = getDate()

// Converting JSON data to string
const data = JSON.stringify({currentDate:date});

fetch("/check", {
    method: "POST", 
    body: data
  }).then(res => {
    return res.text()
  }).then(function(data){
      element.innerHTML=JSON.parse(data).message
  })

function getDate(){
    let yourDate = new Date()
    yourDate.toISOString().split('T')[0]
    const offset = yourDate.getTimezoneOffset()
    yourDate = new Date(yourDate.getTime() - (offset*60*1000))
    return yourDate.toISOString().split('T')[0]
}