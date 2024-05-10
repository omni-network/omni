export const dateFormatter = (date: Date) => {
  if (date instanceof Date === false) {
    return date
  }

  const currentTimestamp = new Date()
  const txsTimestamp = new Date(date)

  const timeDifferenceInMilliseconds = currentTimestamp.getTime() - txsTimestamp.getTime()
  const timeDifferenceInSeconds = Math.abs(timeDifferenceInMilliseconds) / 1000

  // const months = Math.floor(days / 30); // Approximation for simplicity

  // const years = Math.floor(months / 12); // Approximation for simplicity\

  const years = Math.floor(timeDifferenceInSeconds / (3600 * 24 * 30 * 12));
  const months = Math.floor(timeDifferenceInSeconds / (3600 * 24 * 30));
  const days = Math.floor(timeDifferenceInSeconds / (3600 * 24));
  const hours = Math.floor(timeDifferenceInSeconds / 3600)
  const minutes = Math.floor((timeDifferenceInSeconds % 3600) / 60)
  const seconds = Math.floor(timeDifferenceInSeconds % 60)

  // difference greater than a day displays the full date
  if (years > 5) {
    return 'More than 5 years ago'
  }

  // on the second txs show as Just now
  if (years === 0 && months === 0 && days === 0 && hours === 0 && minutes === 0 && seconds === 0) {
    return 'Just now'
  }

  // less than a minute shows seconds ago
  if (years === 0 && months === 0 && days === 0 && hours === 0 && minutes === 0) {
    return `${seconds}s ago`
  }
  // less than an hour, but show minutes
  if (years === 0 && months === 0 && days === 0 && hours === 0) {
    return `${minutes} mins ago`
  }
  if (years === 0 && months === 0 && days === 0) {
    return `${hours} hours ago`
  }
  if (years === 0 && months === 0) {
    return `${days} days ago`
  }
  if (years === 0) {
    return `${months} days ago`
  }
  if (years < 5) {
    return `${years} days ago`
  }
  return txsTimestamp.toLocaleString()
}

export const hashShortener = (hash: string) => {
  if (!hash) {
    return hash
  }
  return `${hash.substring(0, 6)}...${hash.substring(hash.length - 6)}`
}
