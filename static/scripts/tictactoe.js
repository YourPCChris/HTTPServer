var cells = document.querySelectorAll('.cell');
var board = Array(9).fill('');
var currentPlayer = 'X';
var gameOver = false;
var winningCombos = [
    [0, 1, 2], [3, 4, 5], [6, 7, 8],
    [0, 3, 6], [1, 4, 7], [2, 5, 8],
    [0, 4, 8], [2, 4, 6]
];
function checkWin(player) {
    return winningCombos.some(function (combo) { return combo.every(function (index) { return board[index] == player; }); });
}
function isDraw() {
    return board.every(function (cell) { return cell != ''; });
}
function updateGUI() {
    cells.forEach(function (cell, idx) {
        cell.textContent = board[idx];
        cell.disabled = board[idx] !== '' || gameOver;
    });
}
function handleClick(event) {
    if (gameOver)
        return;
    var target = event.target;
    var index = Number(target.dataset.cell);
    if (board[index] !== '')
        return;
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
cells.forEach(function (cell) {
    cell.addEventListener('click', handleClick);
});
updateGUI();
