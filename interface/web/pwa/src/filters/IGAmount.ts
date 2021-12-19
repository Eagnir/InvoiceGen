import Vue from "vue";

Vue.filter('igamount', function (value: any, symbol: string = null, locales: string = null) {
    value = value.toString();
    let style = "decimal";

    if (symbol != null)
        style = "currency";
    if (symbol == null)
        symbol = 'INR';
    if (locales == null)
        locales = 'en-IN';
    const currencyFormater = Intl.NumberFormat(locales, {
        style: style,
        currency: symbol,
        currencyDisplay: "code",
    });
    return "<span class='igamount' data-amount='" + value + "'>" + currencyFormater.format(value).replace(".", '<small>.') + "</small></span>";
})

Vue.filter('igamountRaw', function (value: any, symbol: string = null, locales: string = null) {
    value = value.toString();
    let style = "decimal";

    if (symbol != null)
        style = "currency";
    if (symbol == null)
        symbol = 'INR';
    if (locales == null)
        locales = 'en-IN';
    const currencyFormater = Intl.NumberFormat(locales, {
        style: style,
        currency: symbol,
        currencyDisplay: "code",
    });
    return currencyFormater.format(value);
})