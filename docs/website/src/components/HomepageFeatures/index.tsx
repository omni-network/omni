import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
import Link from '@docusaurus/Link';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: JSX.Element;
  link: string;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'BUILD',
    Svg: require('@site/static/img/Build.svg').default,
    description: (
      <>
      </>
    ),
    link: "/build-xdapp/quickstart",
  },
  {
    title: 'LEARN',
    Svg: require('@site/static/img/Learn.svg').default,
    description: (
      <>
      </>
    ),
    link: "/learn/what-is-omni",
  },
  {
    title: 'FAQ',
    Svg: require('@site/static/img/FAQ.svg').default,
    description: (
      <>
      </>
    ),
    link: "/faq",
  },
];

function Feature({title, Svg, description, link}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <Link href={link}>
      <div className="text--center">
        <Svg className="cta-img" role="img" viewBox="0 0 500 500"/>
      </div>
      <div className="text--center padding-horiz--md">
          <Heading as="h3" className="cta">{title}</Heading>
          <p>{description}</p>
      </div>
      </Link>
    </div>
  );
}

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
