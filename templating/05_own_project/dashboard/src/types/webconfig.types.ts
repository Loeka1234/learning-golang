export interface ComponentOptions {
  Color: string;
  BackColor: string;
}

export interface Component {
  Selected: string;
  Options: ComponentOptions
}

export interface Webconfig {
  Header: Component;
  Section: Component;
  Footer: Component;
}