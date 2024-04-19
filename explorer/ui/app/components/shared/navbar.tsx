import { ThemeButton } from './themebutton'
import logo from '../../../public/Horizontal_Word Black_Symbol Black.png'
import Logo from './logo'
import { Link } from '@remix-run/react'

export default function Navbar() {
  return (
    <header className="static border-none py-3 px-8">
      <div className="w-full flex items-center gap-8">
        <Link to="/" className="m-auto">
          <Logo />
        </Link>

        <div className="grow"></div>

        <Link target="_blank" to="https://forms.gle/EptLH4aFmH5btDWDA" className="group">
          <span className="icon-question text-default text-[20px] group-hover:text-subtle" />{' '}
          <span className="text-default text-b-md group-hover:text-subtle">
            Share your feedback
          </span>
        </Link>
        <span className="m-auto text-default text-b font-normal">EVM Explorer</span>
        <div className="flex">
          <ThemeButton />
        </div>

        <div
          className={`flex gap-2 items-center justify-center border-[1px] border-border-default px-4 py-[14px] rounded-full`}
        >
          <div className="w-3 h-3 rounded-full bg-positive"></div>
          <span className={`text-default `}>Testnet</span>
        </div>
      </div>
    </header>
  )
}
