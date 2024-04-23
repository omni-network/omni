import React, { ButtonHTMLAttributes, ClassAttributes } from 'react'
interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  children: Array<React.ReactNode> | React.ReactNode
  size?: 'sm' | 'md' | 'lg'
  kind?: 'primary' | 'secondary' | 'outline' | 'text'
}

const Button: React.FC<Props> = ({ children, size = 'md', kind = 'primary', ...props }) => {
  const styles = {
    primary: ``,
    secondary: ``,
    outline: ``,
    text: `text-link hover:text-link-hover`,
  }

  return (
    <button {...props} className={`${props.className} ${styles[kind]} text-link`}>
      {children}
    </button>
  )
}

export default Button
