import React from 'react'
interface Props {
  options: Array<any>
  onChange?: Function
  position?: 'left' | 'right' | 'center'
  label?: string
  defaultValue?: string
  placeholder?: string
}

const ChainDropdown: React.FC<Props> = ({
  options,
  onChange = () => {},
  position = 'center',
  label = '',
  defaultValue = '',
  placeholder = 'Filter by',
  ...props
}) => {
  // conditional styles
  const leftStyle = 'rounded-r-none border-r-0'
  const rightStyle = 'rounded-l-none border-l-0'
  const hasLabel = label.length > 0

  // state
  const [isOpen, setIsOpen] = React.useState<boolean>(false)
  const [value, setValue] = React.useState<string>(defaultValue)

  //   const selectedOption =
  React.useEffect(() => {
    onChange(value)
  }, [value])

  return (
    <div className={'relative'}>
      {hasLabel && (
        <label
          className={`absolute z-10 pointer-events-none text-sm text-[12px] font-normal text-subtle left-5 top-1.5 ${isOpen && '!text-default'} ${value !== '' && '!left-[60px]'}`}
        >
          {label}
        </label>
      )}
      <button
        onClick={e => {
          setIsOpen(!isOpen)
        }}
        className={`min-w-[126px] w-full text-nowrap relative text-left px-5 h-[58px] text-cb-md text-subtlest appearance-none rounded-[1000px] bg-[#fcfcfb] bg-opacity-[0.05] border-[1px] border-subtle overflow-hidden ${position === 'left' && leftStyle} ${position === 'right' && rightStyle} ${isOpen && 'bg-overlay bg-opacity-100'} ${value !== '' && 'pl-[70px]'}`}
      >
        {value === '' && (
          <span className={`${isOpen && 'text-default'}  relative top-1.5 block pr-4`}>
            {placeholder}
          </span>
        )}

        {value !== '' && (
          <div className={'flex items-center justify-start gap-2 px-5 pt-3 ml-[-64px]'}>
            {<img src={options.find(option => option.value === value).icon} alt={''} />}
            <span className={'text-cb capitalize text-default pr-5'}>
              {options.find(option => option.value === value).display}
            </span>
          </div>
        )}
      </button>

      {/* chevron */}
      <div className="icon-dropdown-down text-default pointer-events-none absolute text-[20px] top-[20px] right-2" />

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
        <div className="flex flex-col gap-2 absolute top-10 right-0 z-10 bg-overlay border-[1px] border-default rounded-[12px] p-3 min-w-[317px]">
          <input autoFocus className="input h-[40px] rounded-lg bg-search-default" />

          {options.map((option, i) => (
            <div
              key={`option-${i}`}
              onClick={() => {
                setIsOpen(false)
                setValue(option.value)
              }}
              className={`p-2 text-b-md text-default font-bold text-nowrap cursor-pointer rounded-lg flex items-center justify-between ${option.value === value && 'bg-active'} hover:bg-hover`}
            >
              <div className={`flex gap-2`}>
                <img src={option.icon} alt={`${option.display} icon`} />
                <span className='capitalize'>{option.display}</span>
              </div>

              {option.value === value && <span className="icon-tick-med text-[20px]" />}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default ChainDropdown
