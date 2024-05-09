export const dateFormatterXMsgPage = (date) => {
    const options = {
      year: 'numeric',
      month: 'short',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: true,
      timeZoneName: 'short'
    };

    const formatter = new Intl.DateTimeFormat('en-US', options);
    return formatter.format(date);
}
