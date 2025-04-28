class Vechile{
    constructor(brand,model){
        this.brand = brand;
        this.model = model;
    }
    move(){
        console.log("moving ")
    }
}
class Car extends Vechile{
    move(){
        console.log('Drive');
    }
}