import { ThemeButton } from './themebutton'
import logo from '../../../public/Horizontal_Word Black_Symbol Black.png'
import Logo from './logo'
import { Link } from '@remix-run/react'
import ExplorerDropdown from './explorerDropdown'
import Dropdown from './dropdown'
import { useLocation } from '@remix-run/react'
import { useEffect, useState } from 'react'

export default function Navbar() {
  const evmExplorerLinks = [
    {
      value: 'blockScout',
      display: 'on BlockScout',
      icon: 'icon-arrow',
    },
  ]

  const location = useLocation()

  const [currentNet, setCurrentNet] = useState<'testnet' | 'mainnet'>('testnet')

  useEffect(() => {
    console.log(window.location)

    if (window.location.host.includes('test')) {
      setCurrentNet('testnet')
    } else if (window.location.host.includes('main')) {
      setCurrentNet('mainnet')
    }
  }, [])

  return (
    <header className="static border-none py-3 px-8">
      <div className="w-full flex items-center gap-8">
        <Link to="/" className="m-auto">
          <Logo />
        </Link>

        <Dropdown
          onChange={value => {
            if (value === 'testnet') {
              window.location.replace('https://explorer.testnet.omni.network')
            }
            if (value === 'mainnet') {
              window.location.replace('https://explorer.mainnet.omni.network')
            }
          }}
          defaultValue={currentNet}
          options={[
            { display: 'Testnet', value: 'testnet' },
            { display: 'Mainnet', value: 'mainnet' },
          ]}
        />

        <div className="grow"></div>

        <Link target="_blank" to="https://forms.gle/EptLH4aFmH5btDWDA" className="group">
          <span className="icon-question text-default text-[20px] group-hover:text-subtle" />{' '}
          <span className="text-default text-b-md group-hover:text-subtle">
            Share your feedback
          </span>
        </Link>
        <ExplorerDropdown />
        <div className="flex">
          <ThemeButton />
        </div>
      </div>
    </header>
  )
}
