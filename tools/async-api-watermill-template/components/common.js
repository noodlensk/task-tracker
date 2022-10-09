const _ = require('lodash');

export function pascalCase(string) {
    string = _.camelCase(string);
    return string.charAt(0).toUpperCase() + string.slice(1);
}