import React, { useState } from 'react'
interface Props {
  options: Array<string>
  onSelection: any
  value: string
}

interface OptionProps {
  option: string
  active: boolean
  onClick: Function
}

const Option: React.FC<OptionProps> = ({ option, active, onClick }) => {
  return (
    <div
      onClick={() => {
        onClick && onClick()
      }}
      className={`flex gap-1 cursor-pointer px-[18px] py-[7px] text-default font-bold text-center rounded-full border-border-subtle border-[1px] content-center min-w-[70px] bg-bg-input-default hover:border-border-default hover:bg-bg-input-hover ${active && 'bg-bg-input-active border-border-default'}`}
    >
      {active && <span className="icon-tick-med" />}
      {option}
    </div>
  )
}

const FilterOptions: React.FC<Props> = ({ options, onSelection, value, ...props }) => {
  // const [selectedOption, setSelectedOption] = React.useState<string>(options[0])

  // React.useEffect(() => {
  //   onSelection && onSelection(selectedOption)
  // }, [selectedOption])

  return (
    <div {...props} className={`flex gap-2 items-center`}>
      {options.map(option => (
        <Option
          key={option}
          onClick={() => {
            onSelection && onSelection(option)
          }}
          active={value === option}
          option={option}
        />
      ))}
    </div>
  )
}

export default FilterOptions
