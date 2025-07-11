// Game state
var board = Array(9).fill('');
var currentPlayer = 'X';
var gameOver = false;

// DOM elements
var cells = document.querySelectorAll('.cell');
var gameStatus = document.getElementById('gameStatus');
var resetBtn = document.getElementById('reset-btn');

// Winning combinations
var winningCombos = [
    [0, 1, 2], [3, 4, 5], [6, 7, 8], // rows
    [0, 3, 6], [1, 4, 7], [2, 5, 8], // columns
    [0, 4, 8], [2, 4, 6]             // diagonals
];

// Check if a player has won
function checkWin(player) {
    return winningCombos.some(function(combo) {
        return combo.every(function(index) {
            return board[index] === player;
        });
    });
}

// Check if the game is a draw
function isDraw() {
    return board.every(function(cell) {
        return cell !== '';
    });
}

// Update the game board display
function updateBoard() {
    cells.forEach(function(cell, index) {
        cell.textContent = board[index];
        cell.disabled = board[index] !== '' || gameOver;
    });
}

// Update the game status display
function updateStatus() {
    if (gameOver) {
        resetBtn.style.display = 'inline-block';
    } else {
        gameStatus.textContent = 'Current Player: ' + currentPlayer;
        resetBtn.style.display = 'none';
    }
}

// Handle cell clicks
function handleCellClick(event) {
    if (gameOver) return;
    
    var cell = event.target;
    var index = parseInt(cell.getAttribute('data-cell'));
    
    // Check if cell is already filled
    if (board[index] !== '') return;
    
    // Make the move
    board[index] = currentPlayer;
    updateBoard();
    
    // Check for win
    if (checkWin(currentPlayer)) {
        gameStatus.textContent = 'Player ' + currentPlayer + ' Wins!';
        gameOver = true;
        updateStatus();
        return;
    }
    
    // Check for draw
    if (isDraw()) {
        gameStatus.textContent = "It's a Draw!";
        gameOver = true;
        updateStatus();
        return;
    }
    
    // Switch players
    currentPlayer = currentPlayer === 'X' ? 'O' : 'X';
    updateStatus();
}

// Reset the game
function resetGame() {
    board = Array(9).fill('');
    currentPlayer = 'X';
    gameOver = false;
    updateBoard();
    updateStatus();
}

// Add event listeners
cells.forEach(function(cell) {
    cell.addEventListener('click', handleCellClick);
});

resetBtn.addEventListener('click', resetGame);

// Initialize the game
updateBoard();
updateStatus();
