export function parseDate(date: Date): string {
    const strDate = date.toLocaleString().split('T')[0]
    const timeOfDayWithMinutes = date.toLocaleString().split('T')[1].split('.')[0]
    const timeOfDayParts = timeOfDayWithMinutes.split(':')
    const timeOfDay = `${timeOfDayParts[0]}:${timeOfDayParts[1]}`
    return `${strDate}, ${timeOfDay}`
}
