class Vechile{
    constructor(brand,model){
        this.brand = brand;
        this.model = model;
    }
    move(){
        console.log("Moving ....");
    }
}

class Car extends Vechile{
    move(){
        console.log("Drive");
    }
}

class Boat extends Vechile{
    move(){
        console.log('Sail');
    }
}

class Plane extends Vechile{
    move(){
        console.log("Fly!");
    }
}

const car1 = new Car("Ford","Mustang");
const boat1 = new Boat("Ibiza","Touring 20");
const plane1 = new Plane("Boeing","746");

const Vechiles = [car1,boat1,plane1]

Vechiles.forEach(vechile => vechile.move());