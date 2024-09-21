let count = 0;

const value = document.querySelector("#value");
const buttons = document.querySelectorAll(".btn");

buttons.forEach(function (items) {
    items.addEventListener('click', (items) => {
        const targets = items.currentTarget.classList;

        if (targets.contains('decrease')) {
            count--;
        } 
        else if (targets.contains('increase')) {
            count++;
        } 
        else {
            count = 0;
        }
        if (count<0) {
            value.style.color = 'red';
        }
        if (count>0) {
            value.style.color = 'green';
        }
        if (count===0) {
            value.style.color = 'white'
        }
        value.textContent = count;
    })
})