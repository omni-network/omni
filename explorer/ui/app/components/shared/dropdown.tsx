import React from 'react'
interface Props {
  options: Array<{ display: string; value: string }>
  onChange?: Function
  position?: 'left' | 'right' | 'center'
  defaultValue?: string
  isFullWidth?: boolean
}

const Dropdown: React.FC<Props> = ({
  options,
  onChange = () => {},
  position = 'center',
  defaultValue = '',
  isFullWidth = false,
  ...props
}) => {
  // conditional styles
  const leftStyle = 'rounded-r-none border-r-0'
  const rightStyle = 'rounded-l-none border-l-0'

  // state
  const [isOpen, setIsOpen] = React.useState<boolean>(false)
  const [value, setValue] = React.useState<string>(defaultValue)

  return (
    <div className={'relative'}>
      <button
        onClick={e => {
          setIsOpen(!isOpen)
        }}
        className={`flex items-center gap-4 text-left px-3 pl-5 py-2 h-[48px] text-cb-md text-subtlest appearance-none rounded-[1000px]  border-[1px] border-border-subtle overflow-hidden ${position === 'left' && leftStyle} ${position === 'right' && rightStyle} ${isOpen && 'bg-overlay bg-opacity-100 !border-border-default'} ${isFullWidth && 'w-full'}`}
      >
        <span className="text-nowrap text-default">
          {options.find(option => option.value === value)?.display}
        </span>
        <div className='grow'></div>
        <div className="icon-dropdown-down text-default pointer-events-none  text-[20px] " />
      </button>

      {/* dropdown overlay */}
      {isOpen && (
        <div
          onClick={() => {
            setIsOpen(false)
          }}
          className={`fixed w-screen h-screen bg-transparent top-0 left-0 right-0 bottom-0 z-10`}
        />
      )}

      {/* dropdown container */}
      {isOpen && (
        <div className="flex flex-col gap-2 absolute top-10 w-full z-30 bg-overlay border-[1px] border-default rounded-[12px] p-3 min-w-[217px]">
          {options.map((option, i) => (
            <div
              key={`option-${i}`}
              onClick={() => {
                setIsOpen(false)
                setValue(option.value)
                onChange && onChange(option.value)
              }}
              className={`p-2 text-b-md text-default font-bold text-nowrap cursor-pointer rounded-lg flex items-center justify-between ${option.value === value && 'bg-active'} hover:bg-hover`}
            >
              {option.display}
              {option.value === value && <span className="icon-tick-med text-[20px]" />}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default Dropdown
