import React from 'react'

export const BackBtn = ({onBackClickHandler}) => {
  return (
    <button
      onClick={onBackClickHandler}
      type="button"
      className="text-default items-center inline-flex cursor-pointer disabled:cursor-default bg-bg-input-active disabled:bg-bg-input-default rounded-full text-cb-md px-4 py-2 min-w-[60px]"
    >
      <svg
        className="w-6 h-6 me-2"
        aria-hidden="true"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 14 10"
      >
        <path
          stroke="currentColor"
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="1"
          d="M13 5H1m0 0 4 4M1 5l4-4"
        ></path>
      </svg>
      Back
    </button>
  )
}
