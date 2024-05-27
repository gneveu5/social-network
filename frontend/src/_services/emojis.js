const emojiSymbols = [
    {short: ":D", value:"\u{1F600}"},
    {short: "^^", value:"\u{1F601}"},
    {short: "xD", value:"\u{1F923}"},
    {short: ";)", value:"\u{1F609}"},
    {short: "B)", value:"\u{1F60E}"},
    {short: ":)", value:"\u{1F642}"},
    {short: ":|", value:"\u{1F610}"},
    {short: ":)", value:"\u{1F600}"},
    {short: ":o", value:"\u{1F62E}"},
    {short: ":(", value:"\u{1F641}"},
    
]

export const replaceEmojiShortcuts = (text) => {
    let result = text
    emojiSymbols.forEach((elem) => {result = result.replaceAll(elem.short, elem.value)})
    return result
}