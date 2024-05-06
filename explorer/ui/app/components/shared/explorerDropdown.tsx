import React from 'react'
interface Props {}
import BlockScanIcon from '../../assets/images/BlockScan.svg'
import RouteScanIcon from '../../assets/images/RouteScan.svg'

const ExplorerDropdown: React.FC<Props> = ({ ...props }) => {
  // conditional styles
  const leftStyle = 'rounded-r-none border-r-0'
  const rightStyle = 'rounded-l-none border-l-0'

  // state
  const [isOpen, setIsOpen] = React.useState<boolean>(false)

  const items = [
    {
      icon: BlockScanIcon,
      text: 'on Blockscout',
      url: 'https://omni-testnet.blockscout.com',
    },
    {
      icon: RouteScanIcon,
      text: 'on Routescan',
      url: 'https://omni-testnet.routescan.io ',
    },
  ]

  return (
    <div className={'relative'}>
      <button
        onClick={e => {
          setIsOpen(!isOpen)
        }}
        className={`flex gap-2 px-3 flex-row items-center text-nowrap relative text-left h-[48px] text-cb-md text-subtlest appearance-none rounded-[1000px]  border-subtle overflow-hidden ${isOpen && 'bg-overlay bg-opacity-100'} `}
      >
        <label
          className={`z-10 pointer-events-none text-sm text-[12px] font-normal text-subtle  ${isOpen && '!text-default'} `}
        >
          EVM Explorers
        </label>
        {/* chevron */}
        <div
          className={`${isOpen ? 'icon-dropdown-up' : 'icon-dropdown-down'} text-default pointer-events-none text-[20px] `}
        />
      </button>

      {/* dropdown overlay */}
      {isOpen && (
        <div
          onClick={() => {
            setIsOpen(false)
          }}
          className={`fixed w-screen h-screen bg-transparent top-0 left-0 min-[820px]:right-0 bottom-0 z-10`}
        />
      )}
      {/* dropdown container */}
      {isOpen && (
        <div className="flex flex-col gap-2 absolute top-10 min-[820px]:right-0 z-10 bg-overlay border-[1px] border-default rounded-[12px] p-3 ">
          {items.map((option, i) => (
            <a
              href={option.url}
              target="_blank"
              key={`option-${i}`}
              onClick={() => {
                setIsOpen(false)
              }}
              className={`p-2 text-b-md text-default font-bold text-nowrap cursor-pointer rounded-lg flex items-center gap-10 justify-between hover:bg-hover`}
            >
              <div className={`flex gap-2`}>
                <img src={option.icon} alt={`${option.icon} icon`} />
                <span className="">{option.text}</span>
              </div>

              <span className="icon-arrow-right text-[20px]" />
            </a>
          ))}
        </div>
      )}
    </div>
  )
}

export default ExplorerDropdown
