function calculateSums() {
  var sums = [];

  for (var i = 1; i <= 5; i++) {
    var sum = BigInt(0); // Use BigInt to ensure 64-bit integer calculation

    for (var j = 1; j <= 5; j++) {
      if (i !== j) {
        sum += BigInt(j);
      }
    }

    sums.push(sum);
  }

  return sums;
}

var result = calculateSums();
console.log(result); // Output: [14n, 13n, 12n, 11n, 10n]