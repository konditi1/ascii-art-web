
function toggleColors(choice) {
    var colorChoices = document.getElementById('colorChoices');
    if (choice) {
        colorChoices.classList.remove('hidden');
    } else {
        colorChoices.classList.add('hidden');
    }
}