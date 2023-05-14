document.getElementById("prompt").addEventListener("input", function () {
    updateSplitButtonStatus();
    updatePromptCharCount();
});

document.getElementById('split_length').addEventListener('input', updateSplitButtonStatus);

function copyTextToClipboard(text) {
    if (navigator.clipboard) {
        // Modern versions of Chromium browsers, Firefox, etc.
        navigator.clipboard.writeText(text).then(function () {
            console.log("Text successfully copied to clipboard!");
        }, function (error) {
            console.log("Failed to copy text to clipboard: " + error.message);
        });
    } else if (window.clipboardData) {
        // Internet Explorer.
        window.clipboardData.setData("Text", text);

        console.log("Text successfully copied to clipboard!");
    } else {
        // Fallback method using Textarea.

        var textArea = document.createElement("textarea");
        textArea.value = text;
        textArea.style.position = "fixed";
        textArea.style.top = "-999999px";
        textArea.style.left = "-999999px";

        document.body.appendChild(textArea);

        textArea.focus();
        textArea.select();

        try {
            var successful = document.execCommand("copy");

            if (successful) {
                console.log("Text successfully copied to clipboard!");
            } else {
                console.log("Could not copy text to clipboard");
            }
        } catch (error) {
            console.log("Failed to copy text to clipboard: " + error.message);
        }

        document.body.removeChild(textArea);
    }
};

function copyInstructions() {
    const instructionsButton = document.getElementById("copy-instructions-btn");
    const instructions = document.getElementById("instructions").textContent;

    // Copy the instructions to the clipboard.
    navigator.clipboard.writeText(instructions)
        .then(() => {
            console.log('Instructions copied to clipboard');
            instructionsButton.classList.add("clicked");
        })
        .catch(err => {
            console.error('Could not copy text: ', err);
        });
}

function toggleCustomLength(select) {
    const customLengthInput = document.getElementById("split_length");
    if (select.value === "custom") {
        customLengthInput.style.display = "inline";
    } else {
        customLengthInput.value = select.value;
        customLengthInput.style.display = "none";
    }
}

function updateSplitButtonStatus() {
    const promptField = document.getElementById('prompt');
    const splitLength = document.getElementById('split_length');
    const splitBtn = document.getElementById('split-btn');
    const promptLength = promptField.value.trim().length;
    const splitLengthValue = parseInt(splitLength.value);

    if (promptLength === 0) {
        splitBtn.setAttribute('disabled', 'disabled');
        splitBtn.classList.add('disabled');
        splitBtn.textContent = 'Enter a prompt';
    } else if (isNaN(splitLengthValue) || splitLengthValue === 0) {
        splitBtn.setAttribute('disabled', 'disabled');
        splitBtn.classList.add('disabled');
        splitBtn.textContent = 'Enter the length for calculating';
    } else if (promptLength < splitLengthValue) {
        splitBtn.setAttribute('disabled', 'disabled');
        splitBtn.classList.add('disabled');
        splitBtn.textContent = 'Prompt is shorter than split length';
    } else {
        splitBtn.removeAttribute('disabled');
        splitBtn.classList.remove('disabled');
        splitBtn.textContent = `Split into ${Math.ceil(promptLength / splitLengthValue)} parts`;
    }
}

function updatePromptCharCount() {
    const promptField = document.getElementById("prompt");
    const charCount = document.getElementById("prompt-char-count");
    const promptLength = promptField.value.trim().length;
    charCount.textContent = promptLength;
}

updateSplitButtonStatus();
