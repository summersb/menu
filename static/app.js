async function fetchRecipes() {
    const response = await fetch('/api/recipes');
    const recipes = await response.json();
    const recipeList = document.getElementById('recipes-list');
    const menuRecipe = document.getElementById('menu-recipe');
    recipeList.innerHTML = '';
    menuRecipe.innerHTML = '';
    recipes.forEach(recipe => {
        const li = document.createElement('li');
        li.className = 'p-2 border-b';
        li.innerHTML = `<strong>${recipe.name}</strong><br>Ingredients: ${recipe.ingredients.join(', ')}<br>Instructions: ${recipe.instructions}`;
        recipeList.appendChild(li);
        const option = document.createElement('option');
        option.value = recipe.id;
        option.textContent = recipe.name;
        menuRecipe.appendChild(option);
    });

}


async function addRecipe() {
    const name = document.getElementById('recipe-name').value;
    const ingredients = document.getElementById('recipe-ingredients').value.split(',').map(i => i.trim());
    const instructions = document.getElementById('recipe-instructions').value;

    await fetch('/api/recipes', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, ingredients, instructions })
    });
    fetchRecipes();
}


async function addToMenu() {
    const day = document.getElementById('menu-day').value;
    const recipeId = document.getElementById('menu-recipe').value;
    await fetch('/api/menu', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ day_of_week: day, recipe_id: parseInt(recipeId) })
    });
    fetchMenu();
}

async function fetchMenu() {
    const response = await fetch('/api/menu');
    const menu = await response.json();
    const menuList = document.getElementById('menu-list');
    menuList.innerHTML = '';
    menu.forEach(item => {
        const li = document.createElement('li');
        li.className = 'p-2 border-b';
        li.textContent = `${item.day_of_week}: ${item.recipe_name}`;

        menuList.appendChild(li);

    });
}

async function getShoppingList() {
    const response = await fetch('/api/shopping-list');
    const ingredients = await response.json();
    const shoppingList = document.getElementById('shopping-list');

    shoppingList.innerHTML = '';

    ingredients.forEach(ing => {
        const li = document.createElement('li');
        li.className = 'p-2 border-b';
        li.textContent = ing;

        shoppingList.appendChild(li);
    });

}


// Initialize
fetchRecipes();
fetchMenu();
