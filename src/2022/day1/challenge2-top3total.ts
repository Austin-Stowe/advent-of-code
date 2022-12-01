import { readFileSync } from 'fs';

const answer = readFileSync('src/2022/day1/day1-input.txt', 'utf-8')
  .toString()
  .split(/\n\s*\n/)
  .map(elf => elf.split(/\n/))
  .map(elfCalorieList => elfCalorieList
    .map(v => +v)
    .reduce((acc, curr) => acc + curr, 0)
    )
  .sort((a,b) => b - a)
  .slice(0, 3)
  .reduce((acc, curr) => acc + curr, 0);
console.log(answer);