


<script>
    const car = new Set(['yamaha','Toyota','camaro']);
    const list = document.getElementById('carlist');

    car.forEach(brand => {
        const listItem = document.createElement('li');
        listItem.textContent = brand;
        list.appendChild(listItem);
    });
</script>