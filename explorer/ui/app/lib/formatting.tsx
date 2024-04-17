export const dateFormatter = (date: Date) => {
  if (date instanceof Date === false) {
    return date
  }

  const currentTimestamp = new Date()
  const txsTimestamp = new Date(date)

  const timeDifferenceInMilliseconds = currentTimestamp.getTime() - txsTimestamp.getTime()
  const timeDifferenceInSeconds = Math.abs(timeDifferenceInMilliseconds) / 1000

  const hours = Math.floor(timeDifferenceInSeconds / 3600)
  const minutes = Math.floor((timeDifferenceInSeconds % 3600) / 60)
  const seconds = Math.floor(timeDifferenceInSeconds % 60)

  // difference greater than a day displays the full date
  if (hours > 24) {
    return txsTimestamp.toLocaleString()
  }

  // on the second txs show as Just now
  if (hours === 0 && minutes === 0 && seconds === 0) {
    return 'Just now'
  }

  // less than a minute shows seconds ago
  if (hours === 0 && minutes === 0) {
    return `${seconds}s ago`
  }

  // less than an hour, but show minutes
  if (hours === 0 && minutes !== 0) {
    return `${minutes} min, ${seconds}s ago`
  }

  return txsTimestamp.toLocaleString()
}

export const hashShortener = (hash: string) => {
  if (!hash) {
    return hash
  }
  return `${hash.substring(0, 8)}...${hash.substring(hash.length - 5)}`
}
