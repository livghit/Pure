# TemplKit - UI Components

- TemplKit is a CLI tool for generating ready to use components for your project.Please be aware that TemplKit is not an prestyled UI Library , its more like an starting point for your project.

- I look at TemplKit as an tool for generating the most common components that you would use in your project withminimal prestyles. While also having the ability to customize the components to your liking.

# How it will work

1. kit int - Initialize TemplKit in your project
   -This will create an folder named components inside you're project and an config file named templkit.json
   -The config file will contain the path to the components folder and the name of you're project.

2. kit generate [componentname](all components can be found in the official docs) - Generate a component

- This will generate the specified component in the components folder as a .templ file. This file will contain the code for the specified component. From there you can customize the component to your liking.
