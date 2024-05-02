export const copyToClipboard = (value: string) => {
  if (navigator && navigator.clipboard && navigator.clipboard.writeText) {
    return navigator.clipboard.writeText(value)
  }
  return Promise.reject('The Clipboard API is not available.')
}
