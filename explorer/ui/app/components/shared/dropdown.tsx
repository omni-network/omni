import React from 'react'
interface Props {
  options: Array<any>
  onChange?: Function
  position?: 'left' | 'right' | 'center'
  label?: string
  defaultValue?: string
}

const Dropdown: React.FC<Props> = ({
  options,
  onChange = () => {},
  position = 'center',
  label = '',
  defaultValue = '',
  ...props
}) => {
  // conditional styles
  const leftStyle = 'rounded-r-none border-r-0'
  const rightStyle = 'rounded-l-none border-l-0'
  const hasLabel = label.length > 0

  // state
  const [isOpen, setIsOpen] = React.useState<boolean>(false)
  const [value, setValue] = React.useState<string>(defaultValue)

  React.useEffect(() => {
    onChange(value)
  }, [value])

  return (
    <div className={'relative'}>
      {hasLabel && (
        <label className="absolute text-sm text-[12px] font-normal text-subtle left-5 top-2">
          {label}
        </label>
      )}
      <button
        onClick={e => {
          setIsOpen(!isOpen)
        }}
        className={`min-w-[126px] text-left pl-5 h-[58px] text-cb-md text-subtlest appearance-none rounded-[1000px] bg-[#fcfcfb] bg-opacity-[0.05] border-[1px] border-border-subtle overflow-hidden ${position === 'left' && leftStyle} ${position === 'right' && rightStyle} ${isOpen && 'bg-overlay bg-opacity-100 !border-border-default'} `}
      >
        <span className={`${isOpen && 'text-default'}`}>Filter by</span>
      </button>

      {/* chevron */}
      <div className="icon-dropdown-down text-default pointer-events-none absolute text-[20px] top-3.5 right-2" />

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
        <div className="flex flex-col gap-2 absolute top-10 z-10 bg-overlay border-[1px] border-default rounded-[12px] p-3 min-w-[317px]">
          <h5 className="text-b-sm font-bold text-subtle">Filter by:</h5>

          {options.map((option, i) => (
            <div
              key={`option-${i}`}
              onClick={() => {
                setIsOpen(false)
                setValue(option.value)
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
