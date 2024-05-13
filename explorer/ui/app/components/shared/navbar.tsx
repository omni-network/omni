import { ThemeButton } from './themebutton'
import Logo from './logo'
import { Link } from '@remix-run/react'
import ExplorerDropdown from './explorerDropdown'
import Dropdown from './dropdown'
import { useLocation } from '@remix-run/react'
import { useEffect, useState } from 'react'
import { CloseIcon } from '../svg/closeIcon'
import { MenuIcon } from '../svg/menuIcon'

export default function Navbar({ openNav, openNavHandler }) {
  const evmExplorerLinks = [
    {
      value: 'blockScout',
      display: 'on BlockScout',
      icon: 'icon-arrow',
    },
  ]

  const [currentNet, setCurrentNet] = useState<'testnet' | 'mainnet'>('testnet')

  useEffect(() => {
    if (window.location.host.includes('test')) {
      setCurrentNet('testnet')
    } else if (window.location.host.includes('main')) {
      setCurrentNet('mainnet')
    }
  }, [])

  return (
    <header className="static border-none py-3 px-8 relative">
      <div className="w-full flex items-center gap-8 max-[819px]:hidden">
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

        <Link target="_blank" to="https://github.com/orgs/omni-network/discussions" className="group">
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
      <div className={`w-full min-[820px]:hidden`}>
        <nav className={`w-full gap-2 bg-surface flex flex-col ${openNav && 'h-[calc(screen-2rem)] z-20'}`}>
          <div className="flex w-full items-center">
            <Link to="/" className="m-auto relative ">
              <Logo />
            </Link>
            <div className="grow"></div>
            {openNav && <CloseIcon onClick={openNavHandler} />}
            {!openNav && <MenuIcon onClick={openNavHandler} />}
          </div>

          <div className="mt-4"></div>

          {openNav && (
            <>
              <Dropdown
                isFullWidth={true}
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
              <Link target="_blank" to="https://github.com/orgs/omni-network/discussions" className="group">
                <span className="icon-question text-default text-[20px] group-hover:text-subtle" />{' '}
                <span className="text-default text-b-md group-hover:text-subtle">
                  Share your feedback
                </span>
              </Link>
              <ExplorerDropdown />
              <div>
                <ThemeButton />
              </div>
            </>
          )}
        </nav>
      </div>
    </header>
  )
}
