// makes the constant find the button (which i have so graciously id'd)
const button = document.getElementById('bg-toggle-btn');
// listens for clicks
button.addEventListener('click', function() {
    // changes bg, note that this doesn't persist through page refreshes
    document.body.classList.toggle('dark-bg');
});
