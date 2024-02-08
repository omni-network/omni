import React from 'react'
import { classNames } from './utils'

export function Button({ children, className, ...rest }) {
  return (
    <button
      type="button"
      className={classNames(
        'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50',
        className,
      )}
      {...rest}
    >
      {children}
    </button>
  )
}

export function PageButton({ children, className, ...rest }) {
  return (
    <button type="button" className={classNames('btn btn-sm font-medium', className)} {...rest}>
      {children}
    </button>
  )
}
