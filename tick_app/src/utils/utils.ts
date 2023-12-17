const generateURLFriendlyName = (name: string) => {
  return name.toLowerCase().replaceAll(" ", "-");
};

const unTransformURL = (path: string) => {
  return path.replaceAll("-", " ");
};

export { generateURLFriendlyName, unTransformURL };
