const isLetter = (char) => {
    const code = char.charCodeAt(0);
    return (code >= 65 && code <= 90) || (code >= 97 && code <= 122);
};

const isDigit = (char) => {
    const code = char.charCodeAt(0);
    return code >= 48 && code <= 57;
};

module.exports = { isLetter, isDigit };