
const cells = document.querySelectorAll<HTMLButtonElement>('.cell');

type Player = 'X' | '0' | '';
let board: Player[] = Array(9).fill('');
let currentPlayer: Player = 'X';
let gameOver = false;

const winningCombos = [
  [0,1,2], [3,4,5], [6,7,8], 
  [0,3,6], [1,4,7], [2,5,8], 
  [0,4,8], [2,4,6]           
];

function checkWin(player: Player): boolean {
    return winningCombos.some(combo => combo.every(index => board[index] == player));
}

function isDraw(): boolean { 
    return board.every(cell => cell != '');
}

function updateGUI() {
    cells.forEach((cell, idx) => {
        cell.textContent = board[idx];
        cell.disabled = board[idx] !== '' || gameOver;
    });
}

function handleClick(event: MouseEvent) { 
    if (gameOver) return;

    const target = event.target as HTMLButtonElement;
    const index = Number(target.dataset.cell);

    if (board[index] !== '') return;

    board[index] = currentPlayer;
    updateGUI();

    if (checkWin(currentPlayer)) {
        alert('Player ${currentPlayer} Wins!');
        gameOver = true;
        return;
    }

    if (isDraw()) {
        alert("It's a draw!");
        gameOver = true;
        return;
    }

    currentPlayer = currentPlayer === 'X' ? '0' : 'X';
}

cells.forEach(cell => {
    cell.addEventListener('click', handleClick);
});

updateGUI()

