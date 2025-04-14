let i = 1;
while (i <= 3){
    console.log(i);
    i++;
}



for (let k = 0; k < 3; k++){
    console.log(k);
}

for (let x = 0; x < 3; x++){
    console.log("range",x);
}

while (true){
    console.log("loop");
    break
}

for (let n = 0; n < 8784; n++ ){
    if (n % 2 === 0 ){
        continue
    }
    console.log(n)
}