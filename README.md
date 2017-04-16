<h1 align="center">das CMS</h1>

DAS-CMS is a highly customizable and fast CMS for professionals written in Go with a React-based front-end. 


<h2>CURRENTLY IN DEVELOPMENT - NOT READY FOR PRODUCTION USE</h2>


## Why yet another Content Management System?
During my time working with other content-management-systems I found that I was limited.The problem with most CMS is that they're either too heavy for a normal user to use (Typo3, Joomla), or that they lack customization.

**The most popular CMS Wordpress doesn't even support i18n (out of the box).**

What I was looking for was a fast, modular solution for managing the sites of my clients.


### Relieve strain on the end-user

The dream-CMS should allow me as a developer to constrain the System to what exactly needs to be changed by the client (e.g. Content (duh), Imagery, permalinks, users) but not 300 plugins, that might stop working one day, or are incompatible with each-other. **The system HAS to differentiate between a site-administrator and a developer.**

### Be customizable 

As a developer I don't want my clients to have issues managing the website I build for them. I need to it to be simple: So if my users manages a collection of books on their website I don't need a administrative UI managing "articles". I need one that manages "books"! I don't want to fight the template or navigation system. Everything should be declaritive as far as possible and managed through code. 

DAS is not a system targeted at the average user setting up their website. It's supposed to be set up by somebody that knows their stuff and *then* be used by the average user.


###  Be fast

**Most CMS lack optimization:** A website that doesn't change as long as no changes are made to it should *NOT* be generated every time it is requested but rather be generate once and then quickly served. Seeing popular CMS like WordPress and Bolt take more than 500ms to generate a website that isn't changing makes performance oriented developing a nightmare.

Most CMS have still not arrived in the current age of SPAs and API based development. A modern UI that is quick, works on mobile devices and that has a UX that doesn't make the site-administrator want to quit ASAP is a need nowadays.



