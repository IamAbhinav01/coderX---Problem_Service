const marked = require('marked');
const sanitizeHtml = require('sanitize-html');
const TurndownService = require('turndown');
function markdownToHtml(markdown) {
  const turndownService = new TurndownService();

  const convertedMardown = marked.parse(markdown);
  

  const sanitizedHtml = sanitizeHtml(convertedMardown, {
    allowedTags: sanitizeHtml.defaults.allowedTags.concat([
      'img',
      'h1',
      'h2',
      'h3',
      'pre',
      'code',
    ]),
    allowedAttributes: {
      code: ['class'],

      pre: ['class'],

      img: ['src', 'alt'],
    },
  });
  

  const sanitizedMarkdown = turndownService.turndown(sanitizedHtml);
  

  return sanitizedMarkdown;
}

module.exports = markdownToHtml;































































