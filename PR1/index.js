function onSolveButtonClick() {
    let inputA = Number(document.getElementById("inputA").value);
    let inputB = Number(document.getElementById("inputB").value);
    let inputC = Number(document.getElementById("inputC").value);

    let result = document.getElementById("result");
    result.textContent = solveEquation(inputA, inputB, inputC);
}


function solveEquation(a, b, c) {
    if (a === 0) {
        return "Коэффициент a не может быть равен 0!";
    }

    let D = b * b + 4 * a * c;

    if (D < 0) {
        return "Действительных корней нет!";
    }

    else if (D === 0) {
        let x = -b / (2 * a);
        return `Корень уравнения x = ${x}.`;
    }

    else {
        let x1 = (-b + Math.sqrt(D)) / (2 * a);
        let x2 = (-b - Math.sqrt(D)) / (2 * a);

        return `Корни уравнения x1 = ${x1}, x2 = ${x2}.`
    }
}


function onResetButtonClick() {
    document.getElementById("inputA").value = null;
    document.getElementById("inputB").value = null;
    document.getElementById("inputC").value = null;

    document.getElementById("result").textContent = '';
}