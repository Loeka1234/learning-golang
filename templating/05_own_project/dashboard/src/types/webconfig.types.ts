export interface ComponentOptions {
  color: string;
  backColor: string;
}

export interface Component {
  selected: string;
  options: ComponentOptions;
}

export interface Webconfig {
  header: Component;
  section: Component;
  footer: Component;
}
