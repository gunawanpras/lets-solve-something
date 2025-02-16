function main() {

    let arr = [
        [11, 2, 4],
        [4, 5, 6],
        [10, -8, -12]
    ];

    console.log(diagonalDifference(arr));
}

// diagonalDifference calculates the absolute difference between the sums of its diagonals
function diagonalDifference(arr) {
    let dv = [0,0];
    
    for(let i = 0; i < arr.length; i++) {
        dv[0] += parseInt(arr[i][i]);
        dv[1] += parseInt(arr[i][(arr.length-1)-i]);
    }

    // dv.sort();
    let diff = (dv[0]> dv[1]) ? dv[0]- dv[1] : dv[1]- dv[0];

    return diff;
}

main();