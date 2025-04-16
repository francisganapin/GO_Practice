const txt = "The rain in Spain";
const pattern = /^The.*Spain$/;

if (pattern.test(txt)){
    console.log('Yes,we have a match');
} else{
    console.log("No match");
}