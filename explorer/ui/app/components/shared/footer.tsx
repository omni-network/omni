import { Telegram, Discord, Twitter } from './icons'
import 'react-social-icons/discord'
import 'react-social-icons/twitter'
import 'react-social-icons/telegram'
import logo from '../../../public/Horizontal_Word Black_Symbol Black.png'
import { Link } from '@remix-run/react'
import Logo from './logo'

export function Footer() {
  return (
    <footer className="flex flex-col gap-4 py-6 px-8 md:px-20 mt-20 bg-surface">
      <div className={'flex flex-col w-full justify-between sm:flex-row ' }>
        <div className={`flex flex-col gap-4 mb-20 sm:mb-auto`}>
          <Logo />

          <p className="text-b-sm text-subtlest">
            Omni unifies Ethereum’s fragmented layer 2 ecosystem <br /> by establishing a low
            latency and high throughput <br /> global messaging network for all rollups.
          </p>

          <div className={`flex gap-6 mt-4`}>
            <Link target="_blank" to={'https://discord.gg/bKNXmaX9VD'} className={''}>
              <Discord />
            </Link>
            <Link target="_blank" to={'https://twitter.com/OmniFDN'} className={''}>
              <Twitter />
            </Link>
            <Link target="_blank" to={'https://t.me/OmniFDN'} className={''}>
              <Telegram />
            </Link>
          </div>
        </div>

        <div className={'grow'}></div>

        <div className={`flex flex-col gap-4`}>
          <Link to="https://omni.network/">Omni Network</Link>
          <Link to="https://docs.omni.network/">Docs</Link>
          <Link to="https://explorer.omni.network">EVM Explorer</Link>
          <Link to="https://news.omni.network/">Blog</Link>
        </div>
      </div>

      <span className={`text-b-xs text-subtlest block mt-10`}>
        ©2024 The Omni Network | All Rights Reserved
      </span>
    </footer>
  )
}
